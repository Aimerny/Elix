package icu.aimerny.elix.registry;

import icu.aimerny.elix.Elix;
import icu.aimerny.elix.block.ElixBlock;
import icu.aimerny.elix.consts.IdConst;
import net.fabricmc.fabric.api.item.v1.FabricItemSettings;
import net.fabricmc.fabric.api.object.builder.v1.block.FabricBlockSettings;
import net.minecraft.block.Block;
import net.minecraft.item.BlockItem;
import net.minecraft.item.Item;
import net.minecraft.registry.Registries;
import net.minecraft.registry.Registry;
import net.minecraft.util.Identifier;

public class ModBlock {

    public static final Block OBSIDIAN_ORE = new Block(FabricBlockSettings.create().strength(1F));
    public static final ElixBlock ELIX_BLOCK = new ElixBlock(FabricBlockSettings.create().strength(2F));

    public static void register(String id, Block block) {
        Registry.register(Registries.BLOCK, new Identifier(Elix.MOD_ID, id), block);
        Registry.register(Registries.ITEM, new Identifier(Elix.MOD_ID, id), new BlockItem(block, new FabricItemSettings()));
    }

    public static void init() {
        register(IdConst.OBSIDIAN_ORE, OBSIDIAN_ORE);
        register(IdConst.ELIX_BLOCK, ELIX_BLOCK);
    }

}
