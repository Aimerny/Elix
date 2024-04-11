package icu.aimerny.elix.block.entity;

import com.mojang.datafixers.types.Type;
import icu.aimerny.elix.Elix;
import icu.aimerny.elix.consts.IdConst;
import icu.aimerny.elix.inventory.ImplementedInventory;
import icu.aimerny.elix.registry.ModBlock;
import icu.aimerny.elix.registry.ModBlockEntity;
import net.fabricmc.fabric.api.object.builder.v1.block.entity.FabricBlockEntityTypeBuilder;
import net.minecraft.block.BlockState;
import net.minecraft.block.entity.BlockEntity;
import net.minecraft.block.entity.BlockEntityType;
import net.minecraft.inventory.Inventories;
import net.minecraft.item.ItemStack;
import net.minecraft.nbt.NbtCompound;
import net.minecraft.network.listener.ClientPlayPacketListener;
import net.minecraft.network.packet.Packet;
import net.minecraft.network.packet.s2c.play.BlockEntityUpdateS2CPacket;
import net.minecraft.util.collection.DefaultedList;
import net.minecraft.util.math.BlockPos;
import org.jetbrains.annotations.Nullable;

public class ElixBlockEntity extends BlockEntity implements ImplementedInventory {
    public ElixBlockEntity(BlockPos pos, BlockState state) {
        super(ModBlockEntity.ELIX_BLOCK_ENTITY, pos, state);
    }

    private final DefaultedList<ItemStack> items = DefaultedList.ofSize(2, ItemStack.EMPTY);

    private int number = 0;

    @Override
    protected void writeNbt(NbtCompound nbt) {
        nbt.putInt("number", number);
        Inventories.writeNbt(nbt, items);
        super.writeNbt(nbt);
    }

    @Override
    public void readNbt(NbtCompound nbt) {
        super.readNbt(nbt);
        Inventories.readNbt(nbt, items);
        number = nbt.getInt("number");
    }

    @Nullable
    @Override
    public Packet<ClientPlayPacketListener> toUpdatePacket() {
        return BlockEntityUpdateS2CPacket.create(this);
    }

    @Override
    public NbtCompound toInitialChunkDataNbt() {
        return createNbt();
    }

    @Override
    public DefaultedList<ItemStack> getItems() {
        return items;
    }
}
