package icu.aimerny.elix.registry;

import icu.aimerny.elix.Elix;
import icu.aimerny.elix.armor.ObsidianArmor;
import net.fabricmc.fabric.api.itemgroup.v1.FabricItemGroup;
import net.minecraft.item.ItemGroup;
import net.minecraft.item.ItemStack;
import net.minecraft.registry.Registries;
import net.minecraft.registry.Registry;
import net.minecraft.text.Text;
import net.minecraft.util.Identifier;

import java.util.List;

public class ModItemGroup {

    private static final ItemGroup ELIX_MATERIAL_GROUP = FabricItemGroup.builder()
            .icon(() -> new ItemStack(ModItem.OBSIDIAN_INGOT))
            .displayName(Text.translatable("itemGroup.elix.material"))
            .entries((context, entries) -> {
                entries.add(ModItem.OBSIDIAN_INGOT);
                entries.add(ModBlock.OBSIDIAN_ORE.asItem());
                entries.add(ModBlock.ELIX_BLOCK.asItem());
            })
            .build();

    private static final ItemGroup ELIX_ARMOR_GROUP = FabricItemGroup.builder()
            .icon(() -> new ItemStack(ObsidianArmor.OBSIDIAN_MATERIAL_HELMET))
            .displayName(Text.translatable("itemGroup.elix.armor"))
            .entries(((displayContext, entries) -> {
                entries.add(ObsidianArmor.OBSIDIAN_MATERIAL_HELMET);
                entries.add(ObsidianArmor.OBSIDIAN_MATERIAL_CHESTPLATE);
                entries.add(ObsidianArmor.OBSIDIAN_MATERIAL_LEGGINGS);
                entries.add(ObsidianArmor.OBSIDIAN_MATERIAL_BOOTS);
            })).build();

    public static void init(){
        Registry.register(Registries.ITEM_GROUP, new Identifier(Elix.MOD_ID, "material"), ELIX_MATERIAL_GROUP);
        Registry.register(Registries.ITEM_GROUP, new Identifier(Elix.MOD_ID, "armor"), ELIX_ARMOR_GROUP);
    }

}
